import { useLoginUser } from "./useLoginUser";


export const checkSession = () => {
    const { setLoginUser } = useLoginUser();
    const { loginUser } = useLoginUser();
}