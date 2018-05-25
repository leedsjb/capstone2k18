import React from "react";
import { Flex } from "grid-styled";

import Card from "../Card";
import Span from "../Span";
import ColoredAvatar from "../ColoredAvatar";
import Box from "../Box";

const CrewDetailListItem = ({ crewDetail, ...props }) => {
    return (
        <Card {...props}>
            <Flex flexDirection="column" alignItems="center">
                <ColoredAvatar fName={crewDetail.fName} size={48} />
                <Flex mt={3} justifyContent="center">
                    <Span fontWeight="bold" textAlign="center">{`${
                        crewDetail.fName
                    } ${crewDetail.lName}`}</Span>
                </Flex>
                <Box mt={1}>
                    <Span textAlign="center">{crewDetail.position}</Span>
                </Box>
            </Flex>
        </Card>
    );
};

export default CrewDetailListItem;
