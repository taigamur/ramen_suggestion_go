import { memo, VFC } from "react"
import { Button, Drawer, DrawerBody, DrawerContent, DrawerOverlay } from "@chakra-ui/react";


type Props = {
    onClose: () => void;
    isOpen: boolean;
    onClickA: () => void;
    onClickB: () => void;
    onClickC: () => void;
}

export const HeaderMenuDrawer: VFC<Props> = memo((props) => {
    const { onClose, isOpen, onClickA, onClickB, onClickC} = props;

    return (
        <Drawer placement="right" size="xs" onClose={onClose} isOpen={isOpen}>
            <DrawerOverlay>
                <DrawerContent>
                    <DrawerBody p={0} bg="gray.100">
                        <Button w="100%" onClick={onClickA}>
                            Page1
                        </Button>
                        <Button w="100%" onClick={onClickB}>
                            Page2
                        </Button>
                        <Button w="100%" onClick={onClickC}>
                            Page3
                        </Button>
                    </DrawerBody>
                </DrawerContent>
            </DrawerOverlay>
        </Drawer>
    )
})