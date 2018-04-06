import React from "react";

import FlexFullHeight from "./FlexFullHeight";

import Circle from "../Circle";
import Text from "../Text";

const ProfileAvatar = () => {
    return (
        <Circle size={32} p={0}>
            <FlexFullHeight alignItems="center" justifyContent="center">
                <Text>V</Text>
            </FlexFullHeight>
        </Circle>
    );
};

export default ProfileAvatar;
