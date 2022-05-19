import { useContext, useEffect } from "react";
import { useCookies } from "react-cookie";
import { LoginUserContext, LoginUserContextType } from "../providers/LoginUserProvider";

export const useLoginUser = () : LoginUserContextType => {
    const { loginUser } = useContext(LoginUserContext);
    const { setLoginUser} = useContext(LoginUserContext);
    const [cookies, ] = useCookies();

    console.log(cookies)

    useEffect(() => {
        if(!loginUser){
            console.log("cookie : " + cookies.user)
            if(cookies.user){   
                setLoginUser(cookies.user)
            }
        }
    },[loginUser]);

    console.log(loginUser)

    return useContext(LoginUserContext);
}