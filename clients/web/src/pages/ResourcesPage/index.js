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
import Card from "../../components/Card";
import Error from "../../components/Error";
import LoadingSpinner from "../../components/LoadingSpinner";

import { fetchResources } from "../../actions/resources/actions";

class ResourcesPage extends Component {
    componentDidMount() {
        this.props.fetchResources();
    }

    renderResources(resources) {
        if (!resources.pending && resources.data.length > 0) {
            return (
                <Container py={6}>
                    <Flex flexWrap="wrap">
                        {resources.data.map(item => {
                            return (
                                <Box key={item.id} px={2} mt={4} w={1 / 2}>
                                    <a
                                        href={item.link}
                                        rel="noopener noreferrer"
                                        target="_blank"
                                    >
                                        <Card height="100%" p={4}>
                                            <Relative pt="50%" w={1}>
                                                <Absolute
                                                    top={0}
                                                    left={0}
                                                    bottom={0}
                                                    right={0}
                                                >
                                                    <FlexFillHeight
                                                        justifyContent="center"
                                                        alignItems="center"
                                                    >
                                                        <Box w={1 / 2}>
                                                            <Image
                                                                src={
                                                                    item.imageLink
                                                                }
                                                            />
                                                        </Box>
                                                    </FlexFillHeight>
                                                </Absolute>
                                            </Relative>
                                            <Flex
                                                mt={4}
                                                justifyContent="center"
                                            >
                                                <Span
                                                    fontWeight="bold"
                                                    textAlign="center"
                                                >
                                                    {item.name}
                                                </Span>
                                            </Flex>
                                        </Card>
                                    </a>
                                </Box>
                            );
                        })}
                    </Flex>
                </Container>
            );
        }
        return <LoadingSpinner />;
    }

    render() {
        return (
            <FlexFillVH flexDirection="column">
                <Helmet>
                    <title>Resources</title>
                </Helmet>

                <TitleBar title="Resources" />
                {this.props.resources.error ? (
                    <Flex
                        flexDirection="column"
                        alignItems="center"
                        flex={1}
                        justifyContent="center"
                    >
                        <Error />
                    </Flex>
                ) : (
                    <ScrollView>
                        {this.renderResources(this.props.resources)}
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
