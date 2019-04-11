import {getUrl} from "@/api/tools";

export function loginAPI(name, callback, errorCallback) {
  let data = new FormData();
  data.append('name', name)
  window.axios.post(getUrl(""), data, {

  }).then(callback)
    .catch(errorCallback)
}
