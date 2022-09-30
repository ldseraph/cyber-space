import PocketBase, { LocalAuthStore, User, Collection, Record } from 'pocketbase';
import configs from '@/configs';


//  PocketBase.prototype.logout = function(redirect = true) {
//   this.authStore.clear();

//   if (redirect) {
//       replace('/login');
//   }
// };

// Custom auth store to sync the svelte admin store state with the authorized admin instance.
class AppAuthStore extends LocalAuthStore {
  save(token: string, model: User | null) {
    super.save(token, model);

    if (model instanceof User) {
      // setUser(model);
    }
  }
  clear() {
    super.clear();

    // setUser(null);
  }
}

const Client = new PocketBase(
  configs.app.baseUrl,
  configs.app.lang,
  new AppAuthStore('pb_admin_auth')
);

if (Client.authStore.model instanceof User) {
  // setUser(Client.authStore.model);
}

function Get<Q, R>(path: string, queryParams = {} as Q): Promise<R> {
  return Client.send(path, {
    'method': 'GET',
    'params': queryParams,
  })
}

function Post<Q, B, R>(path: string, queryParams = {} as Q, bodyParams = {} as B): Promise<R> {
  // bodyParams = Object.assign({}, bodyParams);
  return Client.send(path, {
    'method': 'POST',
    'params': queryParams,
    'body': bodyParams
  })
}

export {
  Client,
  Get,
  Post,
  User,
  Collection,
  Record
}

export default Client;