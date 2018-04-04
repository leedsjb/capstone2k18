import React, { Component } from "react";

import Tab from "../../components/Tab";

class Tabs extends Component {
    render() {
        return (
            <div>
                {React.Children.map(this.props.children, (child, i) => {
                    return (
                        <Tab
                            active={
                                child.key === this.props.active ? true : false
                            }
                            onClick={() => {
                                this.props.onChange(child.key);
                            }}
                        >
                            {child}
                        </Tab>
                    );
                })}
            </div>
        );
    }
}

export default Tabs;
