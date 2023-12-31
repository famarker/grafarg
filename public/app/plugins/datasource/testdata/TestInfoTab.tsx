// Libraries
import React, { PureComponent } from 'react';

// Types
import { PluginConfigPageProps, DataSourcePluginMeta, DataSourceJsonData } from '@grafarg/data';

interface Props extends PluginConfigPageProps<DataSourcePluginMeta<DataSourceJsonData>> {}

export class TestInfoTab extends PureComponent<Props> {
  constructor(props: Props) {
    super(props);
  }

  render() {
    return (
      <div>
        See github for more information about setting up a reproducible test environment.
        <br />
        <br />
        <a
          className="btn btn-inverse"
          href="https://github.com/famarker/grafarg/tree/master/devenv"
          target="_blank"
          rel="noopener noreferrer"
        >
          GitHub
        </a>
        <br />
      </div>
    );
  }
}
