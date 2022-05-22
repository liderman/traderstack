import Vue from 'vue';
import { format } from 'date-fns';

Vue.filter('datetime', (val) => {
  if (!val) {
    return '';
  }

  if (typeof val == 'object') {
    return format(new Date(val.seconds*1000 + val.nanos/1000), 'yyyy.MM.dd HH:mm:SS');
  }

  return format(val, 'yyyy.MM.dd HH:mm:SS');
});
