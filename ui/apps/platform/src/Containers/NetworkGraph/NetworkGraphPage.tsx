import React, { useState } from 'react';
import {
    PageSection,
    Title,
    Bullseye,
    Spinner,
    Button,
    Divider,
    Toolbar,
    ToolbarContent,
    ToolbarGroup,
    ToolbarItem,
} from '@patternfly/react-core';
import useDeepCompareEffect from 'use-deep-compare-effect';

import { timeWindows } from 'constants/timeWindows';
import useFetchClusters from 'hooks/useFetchClusters';
import useFetchDeploymentCount from 'hooks/useFetchDeploymentCount';
import useURLSearch from 'hooks/useURLSearch';
import { fetchNetworkFlowGraph, fetchNetworkPolicyGraph } from 'services/NetworkService';
import queryService from 'utils/queryService';
import timeWindowToDate from 'utils/timeWindows';
import { isCompleteSearchFilter } from 'utils/searchUtils';

import PageTitle from 'Components/PageTitle';
import useURLParameter from 'hooks/useURLParameter';
import NetworkGraphContainer, { Models } from './NetworkGraphContainer';
import EmptyUnscopedState from './components/EmptyUnscopedState';
import NetworkBreadcrumbs from './components/NetworkBreadcrumbs';
import NetworkSearch from './components/NetworkSearch';
import SimulateNetworkPolicyButton from './simulation/SimulateNetworkPolicyButton';
import EdgeStateSelect, { EdgeState } from './components/EdgeStateSelect';
import DisplayOptionsSelect, { DisplayOption } from './components/DisplayOptionsSelect';
import TimeWindowSelector from './components/TimeWindowSelector';
import {
    transformPolicyData,
    transformActiveData,
    createExtraneousFlowsModel,
    graphModel,
} from './utils/modelUtils';
import { getScopeHierarchyFromSearch } from './utils/hierarchyUtils';
import getSimulation from './utils/getSimulation';

import './NetworkGraphPage.css';
import CIDRFormModal from './components/CIDRFormModal';

const emptyModel = {
    graph: graphModel,
    nodes: [],
    edges: [],
};

// TODO: get real includePorts flag from user input
const includePorts = true;

// for MVP, always show Orchestrator Components
const ALWAYS_SHOW_ORCHESTRATOR_COMPONENTS = true;

function NetworkGraphPage() {
    const [edgeState, setEdgeState] = useState<EdgeState>('active');
    const [displayOptions, setDisplayOptions] = useState<DisplayOption[]>([
        'policyStatusBadge',
        'externalBadge',
        'edgeLabel',
    ]);
    const [models, setModels] = useState<Models>({
        activeModel: emptyModel,
        extraneousModel: emptyModel,
    });
    const [previouslySelectedCluster, setPreviouslySelectedCluster] = useState<string | undefined>(
        undefined
    );

    const [isLoading, setIsLoading] = useState(false);
    const [timeWindow, setTimeWindow] = useState<typeof timeWindows[number]>(timeWindows[0]);
    const [lastUpdatedTime, setLastUpdatedTime] = useState<string>('never');
    const [isCIDRBlockFormOpen, setIsCIDRBlockFormOpen] = useState(false);

    const { searchFilter } = useURLSearch();
    const [simulationQueryValue] = useURLParameter('simulation', undefined);
    const simulation = getSimulation(simulationQueryValue);

    const {
        cluster: clusterFromUrl,
        namespaces: namespacesFromUrl,
        deployments: deploymentsFromUrl,
        remainingQuery,
    } = getScopeHierarchyFromSearch(searchFilter);
    if (clusterFromUrl !== previouslySelectedCluster) {
        setModels({
            activeModel: emptyModel,
            extraneousModel: emptyModel,
        });
        setPreviouslySelectedCluster(clusterFromUrl);
    }

    const hasClusterNamespaceSelected = Boolean(clusterFromUrl && namespacesFromUrl.length);

    const { clusters } = useFetchClusters();
    const selectedClusterId = clusters.find((cl) => cl.name === clusterFromUrl)?.id;
    const selectedCluster = { name: clusterFromUrl, id: selectedClusterId };
    const { deploymentCount } = useFetchDeploymentCount(selectedClusterId || '');

    useDeepCompareEffect(() => {
        // check that user is finished adding a complete filter
        const isQueryFilterComplete = isCompleteSearchFilter(remainingQuery);

        // only refresh the graph data from the API if both a cluster and at least one namespace are selected
        const isClusterNamespaceSelected =
            clusterFromUrl && namespacesFromUrl.length > 0 && deploymentCount;

        if (isQueryFilterComplete && isClusterNamespaceSelected) {
            if (selectedClusterId) {
                setIsLoading(true);

                const queryToUse = queryService.objectToWhereClause(remainingQuery);
                const timestampToUse = timeWindowToDate(timeWindow);

                Promise.all([
                    fetchNetworkFlowGraph(
                        selectedClusterId,
                        namespacesFromUrl,
                        deploymentsFromUrl,
                        queryToUse,
                        timestampToUse || undefined,
                        includePorts,
                        ALWAYS_SHOW_ORCHESTRATOR_COMPONENTS
                    ),
                    fetchNetworkPolicyGraph(
                        selectedClusterId,
                        namespacesFromUrl,
                        deploymentsFromUrl,
                        queryToUse,
                        undefined,
                        includePorts,
                        ALWAYS_SHOW_ORCHESTRATOR_COMPONENTS
                    ),
                ])
                    .then((values) => {
                        // get policy nodes from api response
                        const { nodes: policyNodes } = values[1].response;
                        // transform policy data to DataModel
                        const { policyDataModel, policyNodeMap } = transformPolicyData(policyNodes);
                        // get active nodes from api response
                        const { nodes: activeNodes } = values[0].response;
                        // transform active data to DataModel
                        const { activeDataModel, activeEdgeMap, activeNodeMap } =
                            transformActiveData(activeNodes, policyNodeMap);

                        // create extraneous flows graph
                        const extraneousFlowsDataModel = createExtraneousFlowsModel(
                            policyDataModel,
                            activeNodeMap,
                            activeEdgeMap
                        );

                        const newUpdatedTimestamp = new Date();
                        // show only hours and minutes, use options with the default locale - use an empty array
                        const lastUpdatedDisplayTime = newUpdatedTimestamp.toLocaleTimeString([], {
                            hour: '2-digit',
                            minute: '2-digit',
                        });
                        setLastUpdatedTime(lastUpdatedDisplayTime);

                        setModels({
                            activeModel: activeDataModel,
                            extraneousModel: extraneousFlowsDataModel,
                        });
                    })
                    .catch(() => {
                        // TODO
                    })
                    .finally(() => setIsLoading(false));
            }
        }
    }, [
        clusterFromUrl,
        namespacesFromUrl,
        deploymentsFromUrl,
        remainingQuery,
        timeWindow,
        deploymentCount,
    ]);

    function toggleCIDRBlockForm() {
        setIsCIDRBlockFormOpen(!isCIDRBlockFormOpen);
    }

    return (
        <>
            <PageTitle title="Network Graph (2.0 preview)" />
            <PageSection variant="light" padding={{ default: 'noPadding' }}>
                <Toolbar
                    className="network-graph-selector-bar"
                    data-testid="network-graph-selector-bar"
                >
                    <ToolbarContent>
                        <ToolbarGroup variant="filter-group">
                            <Title headingLevel="h1" className="pf-u-screen-reader">
                                Network Graph
                            </Title>
                            <NetworkBreadcrumbs
                                clusters={clusters}
                                selectedCluster={selectedCluster}
                                selectedNamespaces={namespacesFromUrl}
                                selectedDeployments={deploymentsFromUrl}
                            />
                        </ToolbarGroup>
                        <ToolbarGroup variant="button-group" alignment={{ default: 'alignRight' }}>
                            <ToolbarItem spacer={{ default: 'spacerMd' }}>
                                <Button
                                    variant="secondary"
                                    onClick={toggleCIDRBlockForm}
                                    isDisabled={!selectedClusterId}
                                >
                                    Manage CIDR blocks
                                </Button>
                            </ToolbarItem>
                            <ToolbarItem spacer={{ default: 'spacerNone' }}>
                                <SimulateNetworkPolicyButton
                                    simulation={simulation}
                                    isDisabled={!hasClusterNamespaceSelected}
                                />
                            </ToolbarItem>
                        </ToolbarGroup>
                    </ToolbarContent>
                </Toolbar>
            </PageSection>
            <Divider component="div" />
            <PageSection variant="light" padding={{ default: 'noPadding' }}>
                <Toolbar data-testid="network-graph-toolbar">
                    <ToolbarContent>
                        <ToolbarGroup variant="filter-group">
                            <ToolbarItem>
                                <EdgeStateSelect
                                    edgeState={edgeState}
                                    setEdgeState={setEdgeState}
                                    isDisabled={!hasClusterNamespaceSelected}
                                />
                            </ToolbarItem>
                            <ToolbarItem>
                                <TimeWindowSelector
                                    activeTimeWindow={timeWindow}
                                    setActiveTimeWindow={setTimeWindow}
                                    isDisabled={isLoading || !hasClusterNamespaceSelected}
                                />
                            </ToolbarItem>
                        </ToolbarGroup>
                        <Divider orientation={{ default: 'vertical' }} />
                        <ToolbarGroup className="pf-u-flex-grow-1">
                            <ToolbarItem className="pf-u-flex-grow-1">
                                <NetworkSearch
                                    selectedCluster={clusterFromUrl}
                                    selectedNamespaces={namespacesFromUrl}
                                    selectedDeployments={deploymentsFromUrl}
                                    isDisabled={!hasClusterNamespaceSelected}
                                />
                            </ToolbarItem>
                            <ToolbarItem>
                                <DisplayOptionsSelect
                                    selectedOptions={displayOptions}
                                    setSelectedOptions={setDisplayOptions}
                                    isDisabled={!hasClusterNamespaceSelected}
                                />
                            </ToolbarItem>
                        </ToolbarGroup>
                        <ToolbarGroup alignment={{ default: 'alignRight' }}>
                            <Divider component="div" orientation={{ default: 'vertical' }} />
                            <ToolbarItem className="pf-u-color-200">
                                <em>Last updated at {lastUpdatedTime}</em>
                            </ToolbarItem>
                        </ToolbarGroup>
                    </ToolbarContent>
                </Toolbar>
            </PageSection>
            <Divider component="div" />
            <PageSection
                className="network-graph"
                variant={hasClusterNamespaceSelected ? 'default' : 'light'}
                padding={{ default: 'noPadding' }}
            >
                {!hasClusterNamespaceSelected && <EmptyUnscopedState />}
                {models.activeModel.nodes.length > 0 &&
                    models.extraneousModel.nodes.length > 0 &&
                    !isLoading && (
                        <NetworkGraphContainer
                            models={models}
                            edgeState={edgeState}
                            displayOptions={displayOptions}
                            simulation={simulation}
                            selectedClusterId={selectedClusterId || ''}
                            clusterDeploymentCount={deploymentCount || 0}
                        />
                    )}
                {isLoading && (
                    <Bullseye>
                        <Spinner isSVG />
                    </Bullseye>
                )}
                <CIDRFormModal
                    selectedClusterId={selectedClusterId || ''}
                    isOpen={isCIDRBlockFormOpen}
                    onClose={toggleCIDRBlockForm}
                />
            </PageSection>
        </>
    );
}

export default NetworkGraphPage;
