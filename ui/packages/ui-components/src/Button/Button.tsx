import React, { ReactElement, ReactNode } from 'react';

const buttonClasses =
    'border-2 border-primary-400 font-600 hover:bg-primary-200 hover:text-base-700 inline-flex items-center justify-center min-w-16 px-2 py-2 rounded-sm text-center text-primary-800 text-sm uppercase';

// @TODO This is just starter code for the Button Component. We can discuss, in more detail, how we want to go about it later
/* Maybe omit type prop and separate into 2 components:
 * Button has onClick and children props
 * SubmitButton has children prop
 */
function Button({ type, onClick, children }: ButtonProps): ReactElement {
    if (type === 'submit') {
        return (
            <button className={buttonClasses} type="submit">
                {children}
            </button>
        );
    }
    return (
        <button className={buttonClasses} type="button" onClick={onClick}>
            {children}
        </button>
    );
}

export type ButtonProps = {
    type: 'button' | 'submit';
    onClick: React.MouseEventHandler<HTMLButtonElement>; // required, but not used for type submit
    children: ReactNode;
};

export default Button;
