import React from "react";

import FlexFullHeight from "./FlexFullHeight";

import Span from "../Span";
import Circle from "../Circle";

const ColoredAvatar = ({ fName, onClick, size = 32 }) => {
    return (
        <Circle bg="purple" p={0} onClick={onClick} size={size}>
            <FlexFullHeight alignItems="center" justifyContent="center">
                <Span fontSize={size / 2} fontWeight="normal">
                    {fName.charAt(0)}
                </Span>
            </FlexFullHeight>
        </Circle>
    );
};

export default ColoredAvatar;
