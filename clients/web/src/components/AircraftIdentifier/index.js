import React, { Component } from "react";
import { Flex } from "grid-styled";

import Badge from "../Badge";
import Span from "../Span";
import Box from "../Box";

class AircraftIdentifier extends Component {
    getStatusColor(status) {
        switch (status) {
            case "On a Mission":
                return "#feb204";
            case "Out Of Service":
                return "#e4002b";
            case "Ready":
                return "green";
            default:
                return null;
        }
    }

    render() {
        return (
            <Flex alignItems="center">
                <Span fontSize={this.props.fontSize} fontWeight="bold">
                    {this.props.callsign}
                </Span>
                <Box ml={2}>
                    <Span fontSize={this.props.fontSize}>
                        {this.props.nNum}
                    </Span>
                </Box>
                <Box ml={2}>
                    <Badge statusbg={this.getStatusColor(this.props.status)}>
                        {this.props.status}
                    </Badge>
                </Box>
            </Flex>
        );
    }
}

export default AircraftIdentifier;
