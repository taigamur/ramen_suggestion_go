import {memo, VFC} from "react"
import { useLoginUser } from "../../hooks/useLoginUser";

export const Index: VFC = memo(() => {
    
    const { loginUser } = useLoginUser();

    return(
        <>
            <p>Indexページです。</p>
            <p>{loginUser?.name}さん</p>
        </>

    );
});