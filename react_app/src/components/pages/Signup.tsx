import {memo, VFC, useState, ChangeEvent} from "react"
import { Box, Divider, Flex, Heading, Input, Stack } from "@chakra-ui/react"
import { PrimaryButton} from "../atoms/button/PrimaryButton"
import axios from "axios";
import { User } from "../../types/user" 

export const Signup: VFC = memo(() => {

    const [userName, setUserName] = useState("");
    const onChangeUserName = (e: ChangeEvent<HTMLInputElement>) => setUserName(e.target.value);

    const SignupRequest = () => {
        console.log("test")
        axios.post("http://localhost:3000/signup", {
            mial: "taiga@example.com",
            password: "taiga"
        })
        .then(function(response) {
            console.log(response.data)
        })
    }

    return(
        <Flex align="center" justify="center" height="100vh">
            <Box bg="white" w="sm" p={4} borderRadius="md" shadow="md">
                <Heading as="h1" size="lg" textAlign="center">Sample App</Heading>
                <Divider my={4}/>
                <Stack spacing={6} py={4} px={10}>
                    <Input placeholder="ユーザーネーム" value={userName} onChange={onChangeUserName} />
                    <Input placeholder="パスワード" />
                    <Input placeholder="パスワード(再入力)" />
                    {/* <PrimaryButton disabled={true} loading={true}>ログイン</PrimaryButton> */}
                    <PrimaryButton disabled={userName === ""} onClick={SignupRequest} >登録</PrimaryButton>
                </Stack>

            </Box>
        </Flex>
    )
});