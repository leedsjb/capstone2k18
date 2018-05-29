import sys from "system-components";
import Box from "../Box";

const DropdownItem = sys(
    {
        is: Box,
        bg: "white",
        px: 3,
        py: 2
    },
    props => ({
        "&:hover": {
            backgroundColor: `${props.theme.colors.gray6}`
        }
    })
);

export default DropdownItem;
