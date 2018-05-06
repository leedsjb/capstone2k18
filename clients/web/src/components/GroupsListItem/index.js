import React from "react";
import { Flex } from "grid-styled";

import MasterListItem from "../MasterListItem";
import Span from "../Span";
import Circle from "../Circle";
import Icon from "../Icon";
import Box from "../Box";

const GroupsListItem = ({ group }) => {
    return (
        <MasterListItem>
            <Flex alignItems="center" justifyContent="space-between">
                <Flex alignItems="center">
                    <Circle size={40} />
                    <Box ml={2}>
                        <Span fontWeight="bold">{group.name}</Span>
                        <br />
                        <Span>{group.peoplePreview}</Span>
                    </Box>
                </Flex>
                <Icon glyph="chevronRight" size={16} />
            </Flex>
        </MasterListItem>
    );
};

export default GroupsListItem;