import { useContext, useEffect } from "react";
import { useCookies } from "react-cookie";
import { LoginUserContext, LoginUserContextType } from "../providers/LoginUserProvider";
import { useHistory } from "react-router-dom";

export const useLoginUser = () : LoginUserContextType => {
    const { loginUser } = useContext(LoginUserContext);
    const { setLoginUser} = useContext(LoginUserContext);
    const history = useHistory();
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