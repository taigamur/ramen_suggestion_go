import { memo, VFC } from "react";
import { IconButton } from "@chakra-ui/react";
import { HamburgerIcon } from "@chakra-ui/icons";

type Props = {
    onOpen: () => void;
};

export const HeaderListButton: VFC<Props> = memo((props) => {
    const { onOpen } = props;
    return (
        <IconButton 
            aria-label="menu-button" 
            icon={<HamburgerIcon/>} 
            size="sm" variant="unstyled" 
            display={{base: "block", md: "block"}} 
            onClick={onOpen}
        />
    );
});