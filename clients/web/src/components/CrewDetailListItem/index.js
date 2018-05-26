import React from "react";
import { Flex } from "grid-styled";

import Card from "../Card";
import Span from "../Span";
import ColoredAvatar from "../ColoredAvatar";
import Box from "../Box";

const CrewDetailListItem = ({ crewDetail, ...props }) => {
    return (
        <Card {...props} height="100%">
            <Flex flexDirection="column" alignItems="center" flex={1}>
                <ColoredAvatar fName={crewDetail.fName} size={48} />
                <Flex mt={3} alignItems="center" flexDirection="column">
                    <Box mb={1}>
                        <Span fontWeight="bold" textAlign="center">{`${
                            crewDetail.fName
                        } ${crewDetail.lName}`}</Span>
                    </Box>
                    <Span textAlign="center">{crewDetail.position}</Span>
                </Flex>
            </Flex>
        </Card>
    );
};

export default CrewDetailListItem;
