import { request } from '@umijs/max';
import qs from 'qs';

export function getList(params) {
  return request('/api/admin/article/index', {
    method: 'GET',
    params: params,
    paramsSerializer: function (params) {
      return qs.stringify(params, { arrayFormat: 'brackets' });
    },
  });
}

export function postStore(data) {
  return request('/api/admin/article/create', {
    method: 'POST',
    data: data,
  });
}

export function getData(id) {
  return request('/api/admin/article/show/' + id, {
    method: 'GET',
  });
}

export function postUpdate(id, data) {
  return request('/api/admin/article/edit/' + id, {
    method: 'POST',
    data: data,
  });
}

export function postDelete(id) {
  return request('/api/admin/article/delete/' + id, {
    method: 'POST',
  });
}


export function getCategoryOptions() {
  return request('/api/admin/article-category/options', {
    method: 'GET',
  });
}

export function getTagsOptions(id) {
  return request('/api/admin/article-tag/options', {
    method: 'GET',
  });
}
