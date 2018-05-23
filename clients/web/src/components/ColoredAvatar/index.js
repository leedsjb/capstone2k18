import React from "react";

import FlexFullHeight from "./FlexFullHeight";

import GradientCircle from "../GradientCircle";
import Span from "../Span";

import generateGradient from "../../utils/generateGradient";

const ColoredAvatar = ({ fName, onClick, size = 32 }) => {
    const gradient = generateGradient(fName);

    return (
        <GradientCircle
            p={0}
            onClick={onClick}
            firstcolor={gradient[0]}
            secondcolor={gradient[1]}
            size={size}
        >
            <FlexFullHeight alignItems="center" justifyContent="center">
                <Span fontSize={size / 2} fontWeight="normal">
                    {fName.charAt(0)}
                </Span>
            </FlexFullHeight>
        </GradientCircle>
    );
};

export default ColoredAvatar;
