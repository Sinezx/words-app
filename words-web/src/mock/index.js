import Mock from 'mockjs';

Mock.mock("/api/v1/word/queryword", {
    'total': 5,
    'words': [
        {
            'id|1-100': 1,
            'source_text': '@word',
            'target_text': '@word',
            'rate|1-8.5-10': 1
        },
        {
            'id|1-100': 1,
            'source_text': '@word',
            'target_text': '@word',
            'rate|1-8.5-10': 1
        },
        {
            'id|1-100': 1,
            'source_text': '@word',
            'target_text': '@word',
            'rate|1-8.5-10': 1
        },
        {
            'id|1-100': 1,
            'source_text': '@word',
            'target_text': '@word',
            'rate|1-8.5-10': 1
        },
        {
            'id|1-100': 1,
            'source_text': '@word',
            'target_text': '@word',
            'rate|1-8.5-10': 1
        }
    ]
})

Mock.mock("/api/v1/word/updateword", {
    'message': "word's rate is updated"
})

Mock.mock("/api/v1/word/addword", {
    'id|1-100': 1
})

Mock.mock("/api/v1/reg", {
    'message': "register success"
})

Mock.mock("/api/v1/sayhi", {
    'user_id|1-100': 1
})