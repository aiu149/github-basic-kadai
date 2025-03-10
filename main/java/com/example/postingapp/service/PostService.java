package com.example.postingapp.service;

import java.util.List;
import java.util.Optional;

import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import com.example.postingapp.entity.Post;
import com.example.postingapp.entity.User;
import com.example.postingapp.form.PostRegisterForm;
import com.example.postingapp.repository.PostRepository;

@Service
public class PostService {
    private final PostRepository postRepository;

    public PostService(PostRepository postRepository) {
        this.postRepository = postRepository;
    }

    // 特定のユーザーに紐づく投稿の一覧を作成日時が新しい順で取得する
    public List<Post> findPostsByUserOrderedByCreatedAtDesc(User user) {
        return postRepository.findByUserOrderByCreatedAtDesc(user);
    }
    
    // 指定したidを持つ投稿を取得する
    public Optional<Post> findPostById(Integer id) {
        return postRepository.findById(id);
    }    
    
    @Transactional
    public void createPost(PostRegisterForm postRegisterForm, User user) {
        Post post = new Post();

        post.setTitle(postRegisterForm.getTitle());
        post.setContent(postRegisterForm.getContent());
        post.setUser(user);

        postRepository.save(post);
    }    
}