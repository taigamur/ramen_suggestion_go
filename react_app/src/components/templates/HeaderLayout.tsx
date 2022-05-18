import {memo, VFC, ReactNode} from "react"
import { PrivateHeader } from "../organisms/PrivateHeader";

// 受け取るpropsの型を定義
type Props = {
    children: ReactNode;
}

export const HeaderLayout: VFC<Props> = memo((props) => {
    const { children } = props;
    return (
        <>
            <PrivateHeader/>
            {children}
        </>
    )
});