import {memo, VFC, useState, ChangeEvent} from "react"
import { Box, Divider, Flex, Heading, Input, Stack } from "@chakra-ui/react"
import { PrimaryButton} from "../atoms/button/PrimaryButton"

export const Login: VFC = memo(() => {

    const [userName, setUserName] = useState("");
    const onChangeUserName = (e: ChangeEvent<HTMLInputElement>) => setUserName(e.target.value);

    const LoginRequest = () => {
        console.log("login request test")
    }

    return(
        <Flex align="center" justify="center" height="100vh">
            <Box bg="white" w="sm" p={4} borderRadius="md" shadow="md">
                <Heading as="h1" size="lg" textAlign="center">Ramen Concierge</Heading>
                <Divider my={4}/>
                <Stack spacing={6} py={4} px={10}>
                    <Input placeholder="ユーザーネーム" value={userName} onChange={onChangeUserName}  />
                    <Input placeholder="パスワード" />
                    {/* <PrimaryButton disabled={true} loading={true}>ログイン</PrimaryButton> */}
                    <PrimaryButton disabled={userName === ""} onClick={LoginRequest}>ログイン</PrimaryButton>
                </Stack>

            </Box>
        </Flex>
    )
});