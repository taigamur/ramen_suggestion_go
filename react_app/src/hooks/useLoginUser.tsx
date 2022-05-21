import { useContext, useEffect } from "react";
import { useCookies } from "react-cookie";
import { LoginUserContext, LoginUserContextType } from "../providers/LoginUserProvider";

export const useLoginUser = () : LoginUserContextType => {
    const { loginUser } = useContext(LoginUserContext);
    const { setLoginUser} = useContext(LoginUserContext);
    const [cookies, ] = useCookies();

    useEffect(() => {
        if(!loginUser){
            console.log("cookie : " + cookies.user)
            if(cookies.user){   
                setLoginUser(cookies.user)
            }
        }
    },[loginUser]);

    return useContext(LoginUserContext);
}