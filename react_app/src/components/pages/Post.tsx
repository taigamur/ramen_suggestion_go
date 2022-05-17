import {memo, VFC} from "react"
import { useHistory, useParams } from "react-router-dom";
import { useLoginUser } from "../../hooks/useLoginUser";

export const Post: VFC = memo(() => {

    const history = useHistory()
    const { loginUser } = useLoginUser();

    const username = useParams();

    console.log(username)


    return(
        <>
            {/* <p>{username}のPost</p> */}
            <p>こんにちは、{loginUser?.name}さん</p>

            
        </>
    )
});