import React, { Component } from "react";
import { Helmet } from "react-helmet";
import { Flex } from "grid-styled";
import { connect } from "react-redux";
import { Link } from "react-router-dom";

import FlexFullHeight from "../../components/FlexFullHeight";
import TitleBar from "../../components/TitleBar";
import TabBar from "../../components/TabBar";
import MasterDetailView from "../../components/MasterDetailView";
import MasterView from "../../components/MasterView";
import DetailView from "../../components/DetailView";
import Heading from "../../components/Heading";
import ButtonIcon from "../../components/ButtonIcon";
import NavBar from "../../components/NavBar";
import ProfileAvatar from "../../components/ProfileAvatar";
import MasterListItem from "../../components/MasterListItem";

import { fetchPeople } from "../../actions/people/actions";
import { fetchPeopleDetail } from "../../actions/peopleDetail/actions";

class PeoplePage extends Component {
    componentDidMount() {
        this.props.fetchPeople();
        if (this.props.match.params.id) {
            this.props.fetchPeopleDetail(this.props.match.params.id);
        }
    }

    componentWillReceiveProps(nextProps) {
        if (
            nextProps.match.params.id &&
            !(nextProps.match.params.id === this.props.match.params.id)
        ) {
            this.props.fetchPeopleDetail(nextProps.match.params.id);
        }
    }

    renderMasterView() {
        let content;
        if (!this.props.people.pending && this.props.people.data.length > 0) {
            content = this.props.people.data.map(person => {
                return (
                    <Link to={`/people/${person.id}`}>
                        <MasterListItem key={person.id}>
                            <div>{person.fName}</div>
                        </MasterListItem>
                    </Link>
                );
            });
        }

        return <MasterView>{content}</MasterView>;
    }

    renderDetailView() {
        let content;

        if (
            !this.props.peopleDetail.pending &&
            !Array.isArray(this.props.peopleDetail.data)
        ) {
            let person = this.props.peopleDetail.data;

            content = (
                <Flex flexDirection="column">
                    <ProfileAvatar fName={person.fName} size={72} />
                    <Heading
                        children={`${person.fName} ${person.lName}`}
                        is="h2"
                    />
                    <Heading
                        children={`${person.position}`}
                        is="h3"
                        fontWeight="normal"
                    />
                    <Flex>
                        <ButtonIcon glyph="bubbleChat">Text</ButtonIcon>
                        <ButtonIcon glyph="phone">Call</ButtonIcon>
                        <ButtonIcon glyph="email">Mail</ButtonIcon>
                    </Flex>
                </Flex>
            );
        }

        return <DetailView>{content}</DetailView>;
    }

    render() {
        return (
            <FlexFullHeight flexDirection="column">
                <Helmet>
                    <title>Missions</title>
                </Helmet>

                <TitleBar title="People" />
                <NavBar />

                <MasterDetailView>
                    {this.renderMasterView()}
                    {this.renderDetailView()}
                </MasterDetailView>
                <TabBar />
            </FlexFullHeight>
        );
    }
}

function mapStateToProps(state) {
    return {
        people: state.people,
        peopleDetail: state.peopleDetail
    };
}

const mapDispatchToProps = {
    fetchPeople,
    fetchPeopleDetail
};

export default connect(mapStateToProps, mapDispatchToProps)(PeoplePage);
