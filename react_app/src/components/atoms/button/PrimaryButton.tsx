import { memo, ReactNode, VFC } from "react"
import { Button } from "@chakra-ui/react"

type Props = {
    children: ReactNode;
    loading?: boolean;
    disabled?: boolean;
    onClick: () => void;
};

export const PrimaryButton: VFC<Props> = memo((props) => {
    const { children, disabled = false, loading = false, onClick } = props;
    return (
        <Button bg="teal.400" color="white" _hover={{ opacity: 0.8}} disabled={disabled} isLoading={loading} onClick={onClick}>
            {children}
        </Button>
    )
}) 