import { memo, VFC, useCallback, useContext, useEffect} from "react"
import {Flex, Heading, Box, Button, useDisclosure, Drawer, DrawerBody, DrawerContent, DrawerOverlay} from "@chakra-ui/react"
import { Redirect, useHistory } from "react-router-dom";
import { HeaderListButton } from "../atoms/button/HeaderListButton";
import { LoginUserContext } from "../../providers/LoginUserProvider";
import { useCookies } from "react-cookie";
import { useMessage } from "../../hooks/useMessage";

export const PrivateHeader: VFC = memo(() => {
    const { showMessage } = useMessage()
    const { isOpen, onOpen, onClose} = useDisclosure();

    const history = useHistory();
    const { loginUser } = useContext(LoginUserContext);
    const [cookies, ] = useCookies();

    const onClickHome = useCallback(() => history.push("/home"), []);
    const onClickA = useCallback(() => history.push("/user/info"), []);
    const onClickPlaces = useCallback(() => history.push("/places"), []);
    const onClickC = useCallback(() => {
        const url = "/user/" + loginUser + "/post"
        history.push(url)
    }, []);

    let authentication: boolean = false;
    if(loginUser){
        authentication = true;
    }else if(cookies.user){
        authentication = true;
    }

    console.log(authentication)
    
    if(authentication){
        return (
            <>
                <Flex as="nav" bg="teal.500" color="gray.50" align="center" justify="space-between" padding={{base: 3, md: 5}}>
                    <Flex align="center" as="a" mr={8} _hover={{cursor: "pointer"}} onClick={onClickHome}>
                        <Heading as="h1" fontSize={{base: "md", md:"lg"}}>
                            Ramen Concierge in Tsukuba
                        </Heading>
                    </Flex>
                    <Flex align="center" fontSize="sm" flexGrow={2} display={{base: "none", md:"flex"}}>
                        <Box pr={4}>
                            <Button onClick={onClickA} variant='link' color="white" size="xs">ユーザー情報</Button>
                        </Box>
                        <Box pr={4}>
                            <Button onClick={onClickPlaces} variant='link' color="white" size="xs">MyPlaces</Button>
                        </Box>
                        <Box>
                            <Button onClick={onClickC} variant='link' color="white" size="xs">LinkC</Button>
                        </Box>
                    </Flex>
                    <HeaderListButton onOpen={onOpen}/>
                </Flex>
                <Drawer placement="right" size="xs" onClose={onClose} isOpen={isOpen}>
                    <DrawerOverlay>
                        <DrawerContent>
                            <DrawerBody p={0} bg="gray.100">
                                <Button w="100%" onClick={onClickA}>
                                    ユーザー情報
                                </Button>
                                <Button w="100%" onClick={onClickPlaces}>
                                    Index
                                </Button>
                                <Button w="100%" onClick={onClickC}>
                                    Page3
                                </Button>
                            </DrawerBody>
                        </DrawerContent>
                    </DrawerOverlay>
                </Drawer>
            </>
        );
    }else{
        // showMessage({title: "ログインしてください", status:"error"})
        return <Redirect to="/login"/>;
    }
})