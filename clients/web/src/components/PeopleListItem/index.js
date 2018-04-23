import React from "react";
import { Flex } from "grid-styled";

import MasterListItem from "../../components/MasterListItem";
import Circle from "../../components/Circle";
import Span from "../../components/Span";
import Box from "../../components/Box";

const PeopleListItem = ({ person }) => {
    return (
        <MasterListItem>
            <Flex alignItems="center">
                <Circle size={40} />
                <Box ml={2}>
                    <Span fontWeight="bold">{person.fName}</Span>
                    <br />
                    <Span>{person.position}</Span>
                </Box>
            </Flex>
        </MasterListItem>
    );
};

export default PeopleListItem;
