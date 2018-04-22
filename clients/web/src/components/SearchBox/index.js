import React from "react";
import { Flex } from "grid-styled";

import TextInput from "../TextInput";
import Icon from "../Icon";
import Relative from "../Relative";
import Absolute from "../Absolute";
import FlexFillHeight from "../../components/FlexFillHeight";

const SearchBox = () => {
    return (
        <Relative>
            <TextInput placeholder="Search" pl={4} pr={3} borderRadius={32} />
            <Absolute top={0} bottom={0}>
                <FlexFillHeight alignItems="center" ml={2}>
                    <Icon glyph="search" size={16} />
                </FlexFillHeight>
            </Absolute>
        </Relative>
    );
};

export default SearchBox;
