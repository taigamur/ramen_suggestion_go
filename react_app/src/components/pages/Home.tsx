import {memo, VFC} from "react"
import { useHistory } from "react-router-dom";
import { useLoginUser } from "../../hooks/useLoginUser";

export const Home: VFC = memo(() => {

    const history = useHistory()
    const { loginUser } = useLoginUser();

    if (loginUser === null){
        history.push("/login");
    }

    return(
        <>
            <p>Homeページです。</p>
            <p>こんにちは、{loginUser?.name}さん</p>
        </>
    )
});