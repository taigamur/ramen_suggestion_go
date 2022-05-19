import {memo, VFC} from "react"
import { useCookies } from "react-cookie";
import { useHistory } from "react-router-dom";
import { useLoginUser } from "../../hooks/useLoginUser";
import { useMessage } from "../../hooks/useMessage";
import { PrimaryButton } from "../atoms/button/PrimaryButton";

export const UserInfo: VFC = memo(() => {

    const history = useHistory()
    const { loginUser } = useLoginUser();
    const { setLoginUser } = useLoginUser();

    const [,,removeCookie ] = useCookies();

    const { showMessage } = useMessage();

    const onClick = () => {
        removeCookie("user")
        setLoginUser(null)
        showMessage({title: "ログアウトしました", status:"success"})
        history.push("/login")
    }

    return(
        <>
            <p>ユーザー情報</p>
            

            <PrimaryButton onClick={onClick}>ログアウト</PrimaryButton>
        </>
    )
});