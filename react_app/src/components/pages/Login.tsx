import {memo, VFC, useState, ChangeEvent, useEffect } from "react"
import { Box, Divider, Flex, Heading, Input, Stack } from "@chakra-ui/react"
import { PrimaryButton} from "../atoms/button/PrimaryButton"
import axios from "axios";
import { Redirect, useHistory } from "react-router-dom";
import { useMessage } from "../../hooks/useMessage";
import { useLoginUser } from "../../hooks/useLoginUser";
import { useCookies } from "react-cookie";


export const Login: VFC = memo(() => {
    const history = useHistory()
    const { showMessage } = useMessage()

    const { setLoginUser } = useLoginUser();
    const { loginUser } = useLoginUser();

    const [loading, setLoading] = useState(false)

    const [userName, setUserName] = useState("")
    const [password, setPassword] = useState("")
    const onChangeUserName = (e: ChangeEvent<HTMLInputElement>) => setUserName(e.target.value);
    const onChangePassword = (e: ChangeEvent<HTMLInputElement>) => setPassword(e.target.value);

    const [cookies, setCookie] = useCookies();

    const onClickLoginRequest = () => {
        setLoading(true);
        var params = new URLSearchParams();
        params.append('name',userName)
        params.append('password', password)
        axios.post('http://localhost:8080/login', params)
        .then((res) => {
            if(res.status == 200){
                console.log(res.data.name)
                setLoginUser(res.data.name)
                showMessage({title: "ログインしました", status: "success"});
                setLoading(false);
                setCookie("user",res.data.name)
                
                history.push("/home");
            }
        }).catch((res) => {
            console.log(res)
            showMessage({title: "入力に誤りがあります", status:"error"});
            setLoading(false);
        })
    }

    console.log("loginUser: " + loginUser)

    let authentication: boolean = false;
    if(loginUser){
        authentication = true;
    }else if(cookies.user){
        authentication = true;
    }

    if(authentication){
        return <Redirect to="/home"/>
    }else{
        return(
            <>
                <Flex align="center" justify="center" height="100vh">
                    <Box bg="white" w="sm" p={4} borderRadius="md" shadow="md">
                        <Heading as="h1" size="lg" textAlign="center">Ramen Concierge</Heading>
                        <Divider my={4}/>
                        <Stack spacing={6} py={4} px={10}>
                            <Input placeholder="ユーザーネーム" value={userName} onChange={onChangeUserName}  />
                            <Input placeholder="パスワード" value={password} onChange={onChangePassword}/>
                            {/* <PrimaryButton disabled={true} loading={true}>ログイン</PrimaryButton> */}
                            <PrimaryButton disabled={userName === ""} loading={loading} onClick={onClickLoginRequest}>ログイン</PrimaryButton>
                        </Stack>
                    </Box>
                </Flex> 
            </>
        )
    }
});