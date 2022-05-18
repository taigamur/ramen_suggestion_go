import { createContext, ReactNode, SetStateAction, Dispatch, useState } from "react";

import { User } from "../types/user";

export type LoginUserContextType = {
    loginUser: string | null;
    setLoginUser: Dispatch<SetStateAction<string | null>>
}

export const LoginUserContext = createContext<LoginUserContextType>(
    {} as LoginUserContextType
);



export const LoginUserProvider = (props: {children : ReactNode}) => {
    const { children } = props;
    const [ loginUser, setLoginUser ] = useState<string | null>(null);

    return (
        <LoginUserContext.Provider value={{ loginUser, setLoginUser}}>
            { children }
        </LoginUserContext.Provider>
    ) 
}