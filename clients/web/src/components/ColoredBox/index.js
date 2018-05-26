import React from "react";

import GradientBox from "../GradientBox";

import generateGradient from "../../utils/generateGradient";

const ColoredBox = ({ word, ...props }) => {
    const gradient = generateGradient(word);

    return (
        <GradientBox
            firstcolor={gradient[0]}
            secondcolor={gradient[1]}
            {...props}
        />
    );
};

export default ColoredBox;
