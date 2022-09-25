import { boot } from 'quasar/wrappers';
import { createI18n } from 'vue-i18n';
import { useGlobalStore } from '@/stores/global-store';
import messages from '@/i18n';

export default boot(({ app }) => {
  const global = useGlobalStore();
  const i18n = createI18n({
    locale: global.lang,
    messages,
    missingWarn: false,
    fallbackWarn: false
  });

  // Set i18n instance on app
  app.use(i18n);
});
