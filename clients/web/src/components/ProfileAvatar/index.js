import React from "react";

import FlexFullHeight from "./FlexFullHeight";

import Circle from "../Circle";
import Span from "../Span";

const ProfileAvatar = ({ fName, onClick, size = 32 }) => {
    return (
        <Circle p={0} onClick={onClick} size={size}>
            <FlexFullHeight alignItems="center" justifyContent="center">
                <Span fontSize={size / 2} fontWeight="normal">
                    {fName.charAt(0)}
                </Span>
            </FlexFullHeight>
        </Circle>
    );
};

export default ProfileAvatar;
