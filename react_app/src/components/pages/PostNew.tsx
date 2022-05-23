import { memo, VFC, useState, ChangeEvent, useCallback } from "react"
import { useHistory } from "react-router-dom";
import { useLoginUser } from "../../hooks/useLoginUser";

import { Button, useDisclosure, FormControl, Input, Wrap, WrapItem, Box, Center, Heading  } from '@chakra-ui/react'

import { PostModal } from "../organisms/PostModal";
import axios from "axios";
import { Place } from "../../types/place"
import { useMessage } from "../../hooks/useMessage";
import { PlaceItem } from "../../components/molecules/PlaceItem"
import { useSelectPlace } from "../../hooks/useSelectPlace";

export const PostNew: VFC = memo(() => {

    const history = useHistory()
    const { showMessage } = useMessage();

    const { loginUser } = useLoginUser();
    const [ keyword, setKeyword ] = useState("");
    const onChangeKeyword = (e: ChangeEvent<HTMLInputElement>) => setKeyword(e.target.value);
    const [ places, setPlaces ] = useState<Array<Place>>([]);

    const { isOpen, onOpen, onClose } = useDisclosure()
    const {selectedPlace, onSelectPlace} = useSelectPlace();

    if (loginUser === null){
        history.push("/login");
    }

    const onClickFindPlace = () => {
        if ( keyword == ""){
            showMessage({title: "お店の名前を入力してください", status:"error"})
        }else{
            var params = new URLSearchParams();
            params.append('keyword', keyword)
            axios.post("http://localhost:8080/place/index", params)
            .then((res) => {
                if(res.status === 200){
                    console.log("success")
                    console.log(res.data)
                    setPlaces(res.data)
                }
            }).catch(() => {
                showMessage({title: "キーワードに一致するお店が見つかりません", status:"error"})
            })
        }
    }

    const onClickPlace = useCallback((id: number) => {
        onSelectPlace({id, places, onOpen})
    }, [onOpen, places])


    return(
        <>
            <Box width='80%' mx="auto" pt={5}>
                <FormControl align='center'>
                    <Heading size='sm' pb={2}>投稿するお店を検索</Heading>
                    <Input placeholder='キーワード' value={keyword} onChange={onChangeKeyword}/>
                    <Button onClick={onClickFindPlace} colorScheme='green' variant='outline' mt={1}>検索</Button>
                </FormControl>
            </Box>


            <Wrap pt={10}>
                {places.map((place) => (
                    <WrapItem key={place.id}  width='100%' align='center'>

                        <PlaceItem id={place.id} name={place.name} address={place.address} onClick={onClickPlace}  />
                         
                    </WrapItem>
                ))}
            </Wrap>
            { selectedPlace && 
                <PostModal onClose={onClose} isOpen={isOpen} name={selectedPlace.name} id={selectedPlace.id} address={selectedPlace.address} />  
            }
        </>
    )
});