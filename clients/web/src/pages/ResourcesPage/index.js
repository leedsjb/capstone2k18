import React, { Component } from "react";
import { connect } from "react-redux";
import { Helmet } from "react-helmet";
import { Flex } from "grid-styled";

import TitleBar from "../../components/TitleBar";
import Container from "../../components/Container";
import TabBar from "../../components/TabBar";
import FlexFillVH from "../../components/FlexFillVH";
import ScrollView from "../../components/ScrollView";
import Box from "../../components/Box";
import Image from "../../components/Image";
import Absolute from "../../components/Absolute";
import FlexFillHeight from "../../components/FlexFillHeight";
import Relative from "../../components/Relative";
import Span from "../../components/Span";

import { fetchResources } from "../../actions/resources/actions";

class ResourcesPage extends Component {
    componentDidMount() {
        this.props.fetchResources();
    }

    renderResources(resources) {
        if (!resources.pending && resources.data.length > 0) {
            return (
                <Flex flexWrap="wrap">
                    {resources.data.map(item => {
                        return (
                            <Box w={1 / 2}>
                                <Relative pt="100%" w={1}>
                                    <Absolute
                                        top={0}
                                        left={0}
                                        bottom={0}
                                        right={0}
                                    >
                                        <FlexFillHeight alignItems="center">
                                            <div>
                                                <Image src={item.imageLink} />
                                            </div>
                                        </FlexFillHeight>
                                    </Absolute>
                                </Relative>
                                <Flex mt={2} justifyContent="center">
                                    <Span>{item.name}</Span>
                                </Flex>
                            </Box>
                        );
                    })}
                </Flex>
            );
        }
        return <div>Loading...</div>;
    }

    render() {
        return (
            <FlexFillVH flexDirection="column">
                <Helmet>
                    <title>Resources</title>
                </Helmet>

                <TitleBar title="Resources" />
                {this.props.resources.error ? (
                    <FlexFillVH flexDirection="column">
                        An error has occurred:{" "}
                        {this.props.resources.error.toString()}
                    </FlexFillVH>
                ) : (
                    <ScrollView>
                        <Container>
                            {this.renderResources(this.props.resources)}
                        </Container>
                    </ScrollView>
                )}
                <TabBar />
            </FlexFillVH>
        );
    }
}

function mapStateToProps(state, ownProps) {
    return {
        resources: state.resources
    };
}

const mapDispatchToProps = {
    fetchResources
};

export default connect(mapStateToProps, mapDispatchToProps)(ResourcesPage);
