import { memo, VFC, useCallback} from "react"
import {Flex, Heading, Link, Box, useDisclosure} from "@chakra-ui/react"
import { useHistory } from "react-router-dom";

import { HeaderListButton } from "../atoms/button/HeaderListButton";
import { HeaderMenuDrawer } from "../molecules/HeaderMenuDrawer"


export const Header: VFC = memo(() => {
    const { isOpen, onOpen, onClose} = useDisclosure();

    const history = useHistory();
    const onClickHome = useCallback(() => history.push("/home"), []);

    const onClickA = useCallback(() => history.push("/user/info"), []);
    const onClickB = useCallback(() => history.push("/index"), []);
    const onClickC = useCallback(() => history.push("/index"), []);


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
                        <Link href='/user/info'>ユーザー情報</Link>
                    </Box>
                    <Box pr={4}>
                        <Link href='/index'>index page</Link>
                    </Box>
                    <Link>LinkC</Link>
                </Flex>
                <HeaderListButton onOpen={onOpen}/>
            </Flex>
            <HeaderMenuDrawer onClose={onClose} isOpen={isOpen} onClickA={onClickA} onClickB={onClickB} onClickC={onClickC} />
        </>
    );
})