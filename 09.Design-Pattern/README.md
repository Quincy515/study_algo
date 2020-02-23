# Go 语言设计模式

[设计模式 Golang实现－《研磨设计模式》读书笔记](https://github.com/senghoo/golang-design-pattern)

## Go 语言设计模式的实例代码

### 创建型模式

* [简单工厂模式（Simple Factory）](https://github.com/senghoo/golang-design-pattern/tree/master/00_simple_factory)
* [工厂方法模式（Factory Method）](https://github.com/senghoo/golang-design-pattern/tree/master/04_factory_method): 定义一个用于创建对象的接口，让子类决定实例化哪一个类。工厂方法使一个类的实例化延迟到其子类。
* [抽象工厂模式（Abstract Factory）](https://github.com/senghoo/golang-design-pattern/tree/master/05_abstract_factory): 提供一个创建一系列相关或者相互依赖对象的接口，而无需指定他们具体的类。
* [创建者模式（Builder）](https://github.com/senghoo/golang-design-pattern/tree/master/06_builder): 将一个复杂对象的构建与它表示分离，使得同样的构建构成可以创建不同的表示。
* [原型模式（Prototype）](https://github.com/senghoo/golang-design-pattern/tree/master/07_prototype): 用原型实例指定创建对象的种类，并且通过拷贝这些原型创建新的对象(done浅拷贝、深拷贝)
* [单例模式（Singleton）](https://github.com/senghoo/golang-design-pattern/tree/master/03_singleton): 保证一个类仅有一个实例，并提供一个访问它的全局访问点。

### 结构型模式

* [外观模式（Facade）](https://github.com/senghoo/golang-design-pattern/tree/master/01_facade): 为子系统中的一组接口提供一个一致的界面，此模式定义了一个高层接口，这个接口使得这一子系统更加容易使用。
* [适配器模式（Adapter）](https://github.com/senghoo/golang-design-pattern/tree/master/02_adapter): 将一个类的接口转换成客户端希望的另一个接口。适配器模式使得原本由于接口不兼容而不能一起工作的那些类可以一起工作。
* [代理模式（Proxy）](https://github.com/senghoo/golang-design-pattern/tree/master/09_proxy): 为其他对象提供一中代理，以控制这个对象的访问。
* [组合模式（Composite）](https://github.com/senghoo/golang-design-pattern/tree/master/13_composite): 将对象组合成树形接口，以表示"部分-整体"的层次接口。组合模式使得用户对单个对象和组合对象的使用具有一致性。
* [享元模式（Flyweight）](https://github.com/senghoo/golang-design-pattern/tree/master/18_flyweight): 运用共享技术有效地支持大量细粒度的对象。
* [装饰模式（Decorator）](https://github.com/senghoo/golang-design-pattern/tree/master/20_decorator): 动态地给一个对象添加一些额外的职责，就增加功能来说，装饰模式比生成子类更为灵活。
* [桥模式（Bridge）](https://github.com/senghoo/golang-design-pattern/tree/master/22_bridge): 将抽象化(Abstraction)与实现化(Implementation)脱耦，使得二者可以独立地变化。

### 行为型模式

* [中介者模式（Mediator）](https://github.com/senghoo/golang-design-pattern/tree/master/08_mediator): 用一个中介对象来封装一系列的对象交互。中介这使各对象不需要显示地相互引用，从而使其耦合松散，而且可以独立地改变他们之间的交互。
* [观察者模式（Observer）](https://github.com/senghoo/golang-design-pattern/tree/master/10_observer): 定义了一种一对多的依赖关系，让多个观察者对象同时监听某一个主题对象。这个主题对象在状态发生改变时，会通知所有观察者对象，使他们能够...
* [命令模式（Command）](https://github.com/senghoo/golang-design-pattern/tree/master/11_command): 将一个请求封装为一个对象，从而使你可用不同的请求对客户进行参数化；对请求排队或者记录请求日志，以及支持可撤销的操作。
* [迭代器模式（Iterator）](https://github.com/senghoo/golang-design-pattern/tree/master/12_iterator): 提供一种方法顺序访问一个聚合对象中的各个元素，而又不暴露该对象的内部表示。
* [模板方法模式（Template Method）](https://github.com/senghoo/golang-design-pattern/tree/master/14_template_method): 定义一个操作中的算法的骨架，而将一些具体步骤延迟到子类中。模板方法使得子类可以不改变一个算法的结构即可重定义该算法的某些特定...
* [策略模式（Strategy）](https://github.com/senghoo/golang-design-pattern/tree/master/15_strategy): 它定义了算法家族，分别封装起来，让它们可以相互替换，此模式让算法的变化，不会影响到使用算法的客户。
* [状态模式（State）](https://github.com/senghoo/golang-design-pattern/tree/master/16_state): 当一个对象的内在状态改变时，允许改变其行为，这个对象看起来像是改变了其类。
* [备忘录模式（Memento）](https://github.com/senghoo/golang-design-pattern/tree/master/17_memento): 在不破坏封装性的前提下，捕获一个对象的内部状态，并在该对象之外保存这个状态。这样以后就可以将该对象恢复到原先保存的状态。
* [解释器模式（Interpreter）](https://github.com/senghoo/golang-design-pattern/tree/master/19_interpreter): 给定一个语言，定义它的文法的一种表示，并定义一个解释器，这个解释器使用该表示来解释语言中的句子。
* [职责链模式（Chain of Responsibility）](https://github.com/senghoo/golang-design-pattern/tree/master/21_chain_of_responsibility): 使多个对象都有机会处理请求，从而避免请求的发送者和接收者之间的耦合关系。将这些对象连成一条链，并沿着这条链传递该请求,...
* [访问者模式（Visitor）](https://github.com/senghoo/golang-design-pattern/tree/master/23_visitor): 表示一个作用于某对象结构中的各元素的操作，它使你可以在不改变各元素的类的前提下定义作用于这些元素的新操作。
