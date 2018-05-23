import React from "react";
import { Flex } from "grid-styled";

import MasterListItem from "../../components/MasterListItem";
import ColoredAvatar from "../../components/ColoredAvatar";
import Span from "../../components/Span";
import Box from "../../components/Box";

const PeopleListItem = ({ active, person }) => {
    return (
        <MasterListItem active={active}>
            <Flex alignItems="center">
                <ColoredAvatar fName={person.fName} size={40} />
                <Box ml={3}>
                    <Span fontWeight="bold">{person.fName}</Span>
                    <Box>
                        <Span>{person.position}</Span>
                    </Box>
                </Box>
            </Flex>
        </MasterListItem>
    );
};

export default PeopleListItem;
