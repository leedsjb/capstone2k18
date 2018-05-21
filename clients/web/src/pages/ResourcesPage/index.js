import React, { Component } from "react";
import { connect } from "react-redux";
import { Helmet } from "react-helmet";

import TitleBar from "../../components/TitleBar";
import Container from "../../components/Container";
import TabBar from "../../components/TabBar";
import FlexFillVH from "../../components/FlexFillVH";
import ScrollView from "../../components/ScrollView";

import { fetchResources } from "../../actions/resources/actions";

class ResourcesPage extends Component {
    componentDidMount() {
        this.props.fetchResources();
    }

    renderResources(resources) {
        if (
            !this.props.resources.pending &&
            !Array.isArray(this.props.resources.data)
        ) {
            return <div>Resources</div>;
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
