import {memo, VFC} from "react"
import { useHistory, useLocation, useParams } from "react-router-dom";
import { useLoginUser } from "../../hooks/useLoginUser";


export const Post: VFC = memo((props) => {

    const history = useHistory()
    const { loginUser } = useLoginUser();

    // const username: string = useParams();

    const {username} = useParams<{username: string}>();
    console.log({username})
    return(
        <>
            {username}のPost
        </>
    )
});