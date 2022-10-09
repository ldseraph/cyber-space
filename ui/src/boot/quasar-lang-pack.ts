import { boot } from 'quasar/wrappers'
import { Quasar } from 'quasar'
const langList = import.meta.glob('../../node_modules/quasar/lang/*.mjs')
import { useGlobalStore } from '@/stores/global-store';
import { watch } from 'vue'

function loadLang(lang: string) {
  langList[`../../node_modules/quasar/lang/${lang}.mjs`]().then(lang => {
    Quasar.lang.set(lang.default)
  })
}

export default boot(async () => {
  const global = useGlobalStore();

  try {
    loadLang(global.lang)

    watch(global, val => {
      try {
        loadLang(val.lang)
      } catch (err) {
        console.error(err)
      }
    })

  }
  catch (err) {
    console.error(err)
  }
})

