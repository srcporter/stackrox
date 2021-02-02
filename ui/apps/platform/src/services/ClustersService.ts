import { normalize } from 'normalizr';
import qs from 'qs';

import searchOptionsToQuery from 'services/searchOptionsToQuery';
import { saveFile } from 'services/DownloadService';
import axios from './instance';
import { cluster as clusterSchema } from './schemas';

const clustersUrl = '/v1/clusters';
const clustersEnvUrl = '/v1/clusters-env';
const clusterInitUrl = '/v1/cluster-init';
const upgradesUrl = '/v1/sensorupgrades';
const autoUpgradeConfigUrl = `${upgradesUrl}/config`;
const manualUpgradeUrl = `${upgradesUrl}/cluster`;

export type Cluster = {
    id: string;
    name: string;
};

// @TODO, We may not need this API function after we migrate to a standalone Clusters page
//        Check to see if fetchClusters and fletchClustersByArray can be collapsed
//        into one function
/**
 * Fetches list of registered clusters.
 */
// TODO specify return type after we rewrite without normalize
// eslint-disable-next-line @typescript-eslint/explicit-module-boundary-types
export function fetchClusters() {
    return axios.get<{ clusters: Cluster[] }>(clustersUrl).then((response) => ({
        response: normalize(response.data, { clusters: [clusterSchema] }),
    }));
}

/**
 * Fetches list of registered clusters.
 */
export function fetchClustersAsArray(options?: unknown[]): Promise<Cluster[]> {
    let queryString = '';
    if (options && options.length !== 0) {
        const query = searchOptionsToQuery(options);
        queryString = qs.stringify(
            {
                query,
            },
            {
                addQueryPrefix: true,
                arrayFormat: 'repeat',
                allowDots: true,
            }
        );
    }
    return axios.get<{ clusters: Cluster[] }>(`${clustersUrl}${queryString}`).then((response) => {
        return response?.data?.clusters || [];
    });
}

/**
 * Fetches unwrapped cluster object by ID.
 */
export function getClusterById(id: string): Promise<Cluster | null> {
    return axios.get<{ cluster: Cluster }>(`${clustersUrl}/${id}`).then((response) => {
        return response?.data?.cluster || null;
    });
}

export type AutoUpgradeConfig = {
    enableAutoUpgrade?: boolean;
};

/**
 * Gets the cluster autoupgrade config.
 */
export function getAutoUpgradeConfig(): Promise<AutoUpgradeConfig> {
    return axios.get<{ config: AutoUpgradeConfig }>(autoUpgradeConfigUrl).then((response) => {
        return response?.data?.config || {};
    });
}

/**
 * Saves the cluster autoupgrade config.
 */
export function saveAutoUpgradeConfig(config: AutoUpgradeConfig): Promise<AutoUpgradeConfig> {
    const wrappedObject = { config };
    return axios.post(autoUpgradeConfigUrl, wrappedObject);
}

/**
 * Manually start a sensor upgrade given the cluster ID.
 */
export function upgradeCluster(id: string): Promise<Record<string, never>> {
    return axios.post(`${manualUpgradeUrl}/${id}`);
}

/**
 * Start a cluster cert rotation.
 */
export function rotateClusterCerts(id: string): Promise<Record<string, never>> {
    return axios.post(`${upgradesUrl}/rotateclustercerts/${id}`);
}

/**
 * Manually start a sensor upgrade for an array of clusters.
 */
export function upgradeClusters(ids = []): Promise<Record<string, never>[]> {
    return Promise.all(ids.map((id) => upgradeCluster(id)));
}

/**
 * Fetches cluster by its ID.
 */
// TODO specify return type after we rewrite without normalize
// eslint-disable-next-line @typescript-eslint/explicit-module-boundary-types
export function fetchCluster(id: string) {
    return axios.get(`${clustersUrl}/${id}`).then((response) => ({
        response: normalize(response.data, { cluster: clusterSchema }),
    }));
}

/**
 * Deletes cluster given the cluster ID. Returns an empty object.
 */
export function deleteCluster(id: string): Promise<Record<string, never>> {
    return axios.delete(`${clustersUrl}/${id}`);
}

/**
 * Deletes clusters given a list of cluster IDs.
 */
export function deleteClusters(ids: string[] = []): Promise<Record<string, never>[]> {
    return Promise.all(ids.map((id) => deleteCluster(id)));
}

/**
 * Creates or updates a cluster given the cluster fields.
 */
// TODO specify return type after we rewrite without normalize
// eslint-disable-next-line @typescript-eslint/explicit-module-boundary-types
export function saveCluster(cluster: Cluster) {
    const promise = cluster.id
        ? axios.put(`${clustersUrl}/${cluster.id}`, cluster)
        : axios.post(clustersUrl, cluster);
    return promise.then((response) => ({
        response: normalize(response.data, { cluster: clusterSchema }),
    }));
}

/**
 * Downloads cluster YAML configuration.
 */
export function downloadClusterYaml(id: string, createUpgraderSA = false): Promise<void> {
    return saveFile({
        method: 'post',
        url: '/api/extensions/clusters/zip',
        data: { id, createUpgraderSA },
    });
}

/**
 * Downloads cluster Helm YAML configuration.
 */
export function downloadClusterHelmValuesYaml(id: string): Promise<void> {
    return saveFile({
        method: 'post',
        url: '/api/extensions/clusters/helm-config.yaml',
        data: { id },
    });
}

/**
 * Fetches the KernelSupportAvailable property.
 */
export function fetchKernelSupportAvailable(): Promise<boolean> {
    return axios.get(`${clustersEnvUrl}/kernel-support-available`).then((response) => {
        return Boolean(response?.data?.kernelSupportAvailable);
    });
}

export type InitBundleAttribute = {
    key: string;
    value: string;
};

export type ClusterInitBundle = {
    id?: string;
    name: string;
    createdAt: string;
    createdBy: {
        id: string;
        authProviderId: string;
        attributes: InitBundleAttribute[];
    };
    expiresAt: string;
};

export function fetchClusterInitBundles(): Promise<{ response: { items: ClusterInitBundle[] } }> {
    return axios
        .get<{ items: ClusterInitBundle[] }>(`${clusterInitUrl}/init-bundles`)
        .then((response) => {
            return {
                response: response.data || { items: [] },
            };
        });
}

export function generateClusterInitBundle(
    data: ClusterInitBundle
): Promise<{ response: { meta?: ClusterInitBundle; helmValuesBundle?: string } }> {
    return axios
        .post<{ meta: ClusterInitBundle; helmValuesBundle: string }>(
            `${clusterInitUrl}/init-bundles`,
            data
        )
        .then((response) => {
            return {
                response: response.data || {},
            };
        });
}

export function revokeClusterInitBundles(ids: string[]): Promise<Record<string, never>> {
    return axios.patch(`${clusterInitUrl}/init-bundles/revoke`, { ids });
}
