import { createContext, ReactNode, SetStateAction, Dispatch, useState } from "react";

import { User } from "../types/user";
import { useCookies } from "react-cookie"


type LoginUser = User

export type LoginUserContextType = {
    loginUser: LoginUser | null;
    setLoginUser: Dispatch<SetStateAction<LoginUser | null>>
}

export const LoginUserContext = createContext<LoginUserContextType>(
    {} as LoginUserContextType
);



export const LoginUserProvider = (props: {children : ReactNode}) => {
    const { children } = props;
    const [ loginUser, setLoginUser ] = useState<LoginUser | null>(null);

    return (
        <LoginUserContext.Provider value={{ loginUser, setLoginUser}}>
            { children }
        </LoginUserContext.Provider>
    ) 
}