package cn.lpe234.grpc.grpcdemo.grpcprovider;

import cn.lpe234.grpc.grpcdemo.grpc.UserProviderGrpc;
import cn.lpe234.grpc.grpcdemo.grpc.UserProviderOuterClass;
import io.grpc.stub.StreamObserver;
import lombok.extern.slf4j.Slf4j;
import net.devh.springboot.autoconfigure.grpc.server.GrpcService;

/**
 * Grpc服务暴露
 *
 * @author lpe234
 * @datetime 2018/11/24 14:36
 */
@Slf4j
@GrpcService(UserProviderGrpc.class)
public class UserProvider extends UserProviderGrpc.UserProviderImplBase {

    @Override
    public void getByUserId(UserProviderOuterClass.UserIdRequest request, StreamObserver<UserProviderOuterClass.UserVoReplay> responseObserver) {
        // super.getByUserId(request, responseObserver);

        // 获取请求数据
        long userId = request.getId();
        log.debug("grpc request: userId=" + userId);

        // 构造返回数据
        UserProviderOuterClass.UserVoReplay.Builder userVoReplayBuild = UserProviderOuterClass.UserVoReplay.newBuilder();
        userVoReplayBuild.setId(userId);
        userVoReplayBuild.setUsername("hello world");
        UserProviderOuterClass.UserVoReplay userVoReplay = userVoReplayBuild.build();

        // 做出响应
        responseObserver.onNext(userVoReplay);
        responseObserver.onCompleted();
    }
}
