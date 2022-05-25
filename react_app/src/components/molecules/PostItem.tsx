import { memo, VFC } from "react"
import { Box, Stack, Text, Button, IconButton, Spacer, Flex, Center, Circle } from "@chakra-ui/react"
import { DeleteIcon, EditIcon } from "@chakra-ui/icons"


import { Post } from "../../types/post"

type Props = {
    post: Post;
}
export const PostItem: VFC<Props> = memo((props) => {
    const { post } = props;


    return(
        <Box w="80%" mx="auto" borderWidth='3px' borderRadius='lg'>
            <Flex>
            <Box w='30%' align="center"  my="auto" >
                <Circle size='60px' bg='green.300'>
                    {post.date.substr(5,5)}
                </Circle>
            </Box>
            <Spacer/>
            <Box w="60%" h='auto'>
                <Text>user: {post.username}</Text>
                <Text>place: {post.place.id}</Text>
                <Text>place name: {post.place.name}</Text>
                <Text>point: {post.value}</Text>  
            </Box>
            <Box w="10%" align="center" >
                <IconButton size='xs' variant='outline' colorScheme='teal' aria-label='Call Sage' icon={<DeleteIcon />} />
            </Box>
            </Flex>
        </Box>
    )
})