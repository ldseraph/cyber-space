import { boot } from 'quasar/wrappers';
import { createI18n,I18nOptions  } from 'vue-i18n';
import { useGlobalStore } from '@/stores/global-store';
import { messages, datetimeFormats } from '@/i18n';

export default boot(({ app }) => {
  const global = useGlobalStore();
  const i18n = createI18n<I18nOptions>({
    locale: global.lang,
    messages,
    datetimeFormats,
    missingWarn: false,
    fallbackWarn: false
  });

  // Set i18n instance on app
  app.use(i18n);
});
