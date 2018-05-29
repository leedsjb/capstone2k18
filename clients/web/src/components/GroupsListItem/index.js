import React from "react";
import { Flex } from "grid-styled";

import MasterListItem from "../MasterListItem";
import Span from "../Span";
import Box from "../Box";
import ColoredAvatar from "../ColoredAvatar";

const GroupsListItem = ({ active, group }) => {
    return (
        <MasterListItem active={active}>
            <Flex alignItems="center" justifyContent="space-between">
                <Flex alignItems="center">
                    <ColoredAvatar size={40} fName={group.name} />
                    <Box ml={3}>
                        <Span fontWeight="bold">{group.name}</Span>
                        <br />
                        <Span>{group.peoplePreview}</Span>
                    </Box>
                </Flex>
            </Flex>
        </MasterListItem>
    );
};

export default GroupsListItem;
