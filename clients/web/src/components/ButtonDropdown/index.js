import React from "react";
import Button from "../Button";

const ButtonDropdown = ({ children, ...props }) => {
    return (
        <Button {...props}>
            {children}
        </Button>
    );
};

export default ButtonDropdown;