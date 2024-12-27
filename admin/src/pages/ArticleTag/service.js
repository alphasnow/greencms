import { request } from '@umijs/max';
import qs from 'qs';

export function getList(params) {
  return request('/api/admin/article-tag/index', {
    method: 'GET',
    params: params,
    paramsSerializer: function (params) {
      return qs.stringify(params, { arrayFormat: 'brackets' });
    },
  });
}

export function postStore(data) {
  return request('/api/admin/article-tag/create', {
    method: 'POST',
    data: data,
  });
}

export function getData(id) {
  return request('/api/admin/article-tag/show/' + id, {
    method: 'GET',
  });
}

export function postUpdate(id, data) {
  return request('/api/admin/article-tag/edit/' + id, {
    method: 'POST',
    data: data,
  });
}

export function postDelete(id) {
  return request('/api/admin/article-tag/delete/' + id, {
    method: 'POST',
  });
}
