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
import Tab from "../../components/Tab";
import Divider from "../../components/Divider";
import Icon from "../../components/Icon";

import { fetchPeople } from "../../actions/people/actions";
import { fetchPeopleDetail } from "../../actions/peopleDetail/actions";

import { fetchGroups } from "../../actions/groups/actions";
import { fetchGroupsDetail } from "../../actions/groupsDetail/actions";

import matchPath from "../../utils/matchPath";

class PeoplePage extends Component {
    componentDidMount() {
        this.props.fetchPeople();
        this.props.fetchGroups();

        if (this.props.match.params.id) {
            this.props.fetchPeopleDetail(this.props.match.params.id);
        }

        if (this.props.match.params.groupId) {
            this.props.fetchGroupsDetail(this.props.match.params.groupId);
        }
    }

    componentWillReceiveProps(nextProps) {
        if (
            nextProps.match.params.id &&
            !(nextProps.match.params.id === this.props.match.params.id)
        ) {
            this.props.fetchPeopleDetail(nextProps.match.params.id);
        }

        if (
            nextProps.match.params.groupId &&
            !(
                nextProps.match.params.groupId ===
                this.props.match.params.groupId
            )
        ) {
            this.props.fetchGroupsDetail(nextProps.match.params.groupId);
        }
    }

    isPeopleTab() {
        return matchPath(this.props.location.pathname, "/people");
    }

    isGroupDetailView() {
        return matchPath(this.props.location.pathname, "/groups/:groupId");
    }

    renderPeopleList() {
        if (!this.props.people.pending && this.props.people.data.length > 0) {
            return this.props.people.data.map(person => {
                return (
                    <Link to={`/people/${person.id}`} key={person.id}>
                        <MasterListItem>
                            <div>{person.fName}</div>
                        </MasterListItem>
                    </Link>
                );
            });
        }
    }

    renderGroupsList() {
        if (!this.props.groups.pending && this.props.groups.data.length > 0) {
            return this.props.groups.data.map(group => {
                return (
                    <Link to={`/groups/${group.id}`} key={group.id}>
                        <MasterListItem>
                            <div>{group.name}</div>
                            <div>{group.peoplePreview}</div>
                        </MasterListItem>
                    </Link>
                );
            });
        }
    }

    renderGroupsDetailList() {
        if (
            !this.props.groupsDetail.pending &&
            !Array.isArray(this.props.groupsDetail.data)
        ) {
            const { id } = this.props.groupsDetail.data;

            return this.props.groupsDetail.data.people.map(person => {
                return (
                    <Link to={`/groups/${id}/${person.id}`} key={person.id}>
                        <MasterListItem>
                            <div>{person.fName}</div>
                        </MasterListItem>
                    </Link>
                );
            });
        }
    }

    renderMasterView() {
        let list;
        if (this.isPeopleTab()) {
            list = this.renderPeopleList();
        } else if (!this.isGroupDetailView()) {
            list = this.renderGroupsList();
        } else {
            list = this.renderGroupsDetailList();
        }

        let controller = !this.isGroupDetailView() ? (
            <Flex>
                <Tab active={this.isPeopleTab()} is={Link} to="/people">
                    People
                </Tab>
                <Tab active={!this.isPeopleTab()} is={Link} to="/groups">
                    Groups
                </Tab>
            </Flex>
        ) : (
            <div>
                <Link to="/groups">
                    <Icon glyph="chevronLeft" size={16} />
                </Link>
                {this.props.groupsDetail.data.name}
            </div>
        );

        return (
            <MasterView>
                {controller}
                <Divider />
                {list}
            </MasterView>
        );
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
        peopleDetail: state.peopleDetail,
        groups: state.groups,
        groupsDetail: state.groupsDetail,
        location: state.router.location
    };
}

const mapDispatchToProps = {
    fetchPeople,
    fetchPeopleDetail,
    fetchGroups,
    fetchGroupsDetail
};

export default connect(mapStateToProps, mapDispatchToProps)(PeoplePage);
