import { memo, ReactNode, VFC } from "react"
import { Box, Stack, Text, Image  } from "@chakra-ui/react"

type Props = {
    id: number;
    name: string;
    onClick: (id: number) => void;
};

export const UserCard: VFC<Props> = memo((props) => {
    const {  id, name, onClick } = props;

    return (
        <Box w="260px" bg="white" borderRadius="10px" shadow="md" p={4} _hover={{ cursor: "pointer", opacity: 0.8}} onClick={() => onClick(id)}>
            <Stack textAlign="center">
            <Text fontSize="lg" fontWeight="bold">{id}</Text>
                <Text fontSize="lg" fontWeight="bold">{name}</Text>
            </Stack>
        </Box>
    )
}) 