import 'jquery';
import '../../../public/vendor/flot/jquery.flot.js';
import '../../../public/vendor/flot/jquery.flot.selection';
import '../../../public/vendor/flot/jquery.flot.time';
import '../../../public/vendor/flot/jquery.flot.stack';
import '../../../public/vendor/flot/jquery.flot.stackpercent';
import '../../../public/vendor/flot/jquery.flot.fillbelow';
import '../../../public/vendor/flot/jquery.flot.crosshair';
import '../../../public/vendor/flot/jquery.flot.dashes';
import '../../../public/vendor/flot/jquery.flot.gauge';
import { withTheme } from '../src/utils/storybook/withTheme';
import { withPaddedStory } from '../src/utils/storybook/withPaddedStory';
// @ts-ignore
import lightTheme from '../../../public/sass/grafarg.light.scss';
// @ts-ignore
import darkTheme from '../../../public/sass/grafarg.dark.scss';
import { GrafargLight, GrafargDark } from './storybookTheme';
import addons from '@storybook/addons';

const handleThemeChange = (theme: any) => {
  if (theme !== 'light') {
    lightTheme.unuse();
    darkTheme.use();
  } else {
    darkTheme.unuse();
    lightTheme.use();
  }
};

addons.setConfig({
  showRoots: false,
  theme: GrafargDark,
});

export const decorators = [withTheme(handleThemeChange), withPaddedStory];

export const parameters = {
  info: {},
  docs: {
    theme: GrafargDark,
  },
  darkMode: {
    dark: GrafargDark,
    light: GrafargLight,
  },
  options: {
    showPanel: true,
    panelPosition: 'right',
    showNav: true,
    isFullscreen: false,
    isToolshown: true,
    storySort: {
      method: 'alphabetical',
      // Order Docs Overview and Docs Overview/Intro story first
      order: ['Docs Overview', ['Intro']],
    },
  },
  knobs: {
    escapeHTML: false,
  },
};
